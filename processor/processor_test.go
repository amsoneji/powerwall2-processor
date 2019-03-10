package main

import (
  "testing"
  "gotest.tools/assert"
  "reflect"
)

func TestPrepareNilJsonObject(t *testing.T){
  output_object := prepareJsonObject("does not exist")
  assert.Assert(t, len(output_object.Struct) == 0)
  assert.Assert(t, len(output_object.List) == 0)
}

func TestPrepareAggJsonObject(t *testing.T){
  output_object := prepareJsonObject("/api/meters/aggregates")
  assert.Assert(t, len(output_object.Struct) != 0)
  assert.Assert(t, len(output_object.List) == 0)
}

func TestPrepareSiteJsonObject(t *testing.T){
  output_object := prepareJsonObject("/api/meters/site")
  assert.Assert(t, len(output_object.Struct) == 0)
  assert.Assert(t, len(output_object.List) != 0)
}

func TestPrepareSolarJsonObject(t *testing.T){
  output_object := prepareJsonObject("/api/meters/solar")
  assert.Assert(t, len(output_object.Struct) == 0)
  assert.Assert(t, len(output_object.List) != 0)
}

func TestFlattenJson(t *testing.T){
  test_input := map[string]interface{}{
    "test1": "value1",
    "test2": map[string]interface{}{
      "test2.1": "value2.1",
      "test2.2": map[string]interface{}{
        "test2.2.1": "value2.2.1",
      },
    },
  }
  obj, err := flattenJson(test_input)
  assert.Assert(t, err == nil)
  expected_output := map[string]interface{}{
    "test1": "value1",
    "test2.test2.1": "value2.1",
    "test2.test2.2.test2.2.1": "value2.2.1",
  }
  assert.Assert(t, reflect.DeepEqual(obj, expected_output))
}

