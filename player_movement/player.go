components {
  id: "player_movement"
  component: "/player_movement/player_movement.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "spaceship_sprite"
  type: "sprite"
  data: "tile_set: \"/assets/sprites/spaceship/space_shooter.atlas\"\n"
  "default_animation: \"spaceship\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 1.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  scale {
    x: 4.0
    y: 4.0
    z: 1.0
  }
}
