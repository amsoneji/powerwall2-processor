package main

import (
  "testing"
  "gotest.tools/assert"
)

func TestGetAllPaths(t *testing.T){
  paths := get_all_paths()
  assert.Equal(t, 3, len(paths))
  assert.Equal(t, "/api/meters/aggregates", paths[0])
  assert.Equal(t, "/api/meters/site", paths[1])
  assert.Equal(t, "/api/meters/solar", paths[2])
}

func TestGetObjForPath(t *testing.T){
  var struct_return map[string]interface{} 
  var list_return []map[string]interface{}

  // Test nil
  struct_return, list_return = get_obj_for_path("does not exist")
  assert.Assert(t, len(struct_return) == 0 )
  assert.Assert(t, len(list_return) == 0)

  // Test "/api/meters/aggregates"
  struct_return, list_return = get_obj_for_path("/api/meters/aggregates")
  assert.Assert(t, len(struct_return) != 0)
  assert.Assert(t, len(list_return) == 0)

  // Test "/api/meters/site"
  struct_return, list_return = get_obj_for_path("/api/meters/site")
  assert.Assert(t, len(struct_return) == 0)
  assert.Assert(t, len(list_return) != 0)

  // Test "/api/meters/solar"
  struct_return, list_return = get_obj_for_path("/api/meters/solar")
  assert.Assert(t, len(struct_return) == 0)
  assert.Assert(t, len(list_return) != 0)
  
}