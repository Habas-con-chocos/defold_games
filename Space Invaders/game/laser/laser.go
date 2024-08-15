components {
  id: "laser"
  component: "/game/laser/laser.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"laser\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/sprite_atlas.atlas\"\n"
  "}\n"
  ""
}
