components {
  id: "mainmenu"
  component: "/main/mainmenu/mainmenu.gui"
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
  id: "Background !!"
  type: "sprite"
  data: "tile_set: \"/main/background/backgroundsprites.atlas\"\n"
  "default_animation: \"Background !!\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 525.0
    y: 262.5
    z: 0.8
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "cloud1"
  type: "sprite"
  data: "tile_set: \"/main/background/backgroundsprites.atlas\"\n"
  "default_animation: \"cloud 1\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 108.0
    y: 278.0
    z: 0.9
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "cloud2"
  type: "sprite"
  data: "tile_set: \"/main/background/backgroundsprites.atlas\"\n"
  "default_animation: \"cloud 2\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 303.0
    y: 432.0
    z: 0.9
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "cloud3"
  type: "sprite"
  data: "tile_set: \"/main/background/backgroundsprites.atlas\"\n"
  "default_animation: \"cloud 3\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 492.0
    y: 272.0
    z: 0.9
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "cloud4"
  type: "sprite"
  data: "tile_set: \"/main/background/backgroundsprites.atlas\"\n"
  "default_animation: \"cloud 4\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 765.0
    y: 427.0
    z: 0.9
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "cloud5"
  type: "sprite"
  data: "tile_set: \"/main/background/backgroundsprites.atlas\"\n"
  "default_animation: \"cloud 5\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 948.0
    y: 240.0
    z: 0.9
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "title"
  type: "sprite"
  data: "tile_set: \"/main/gameover/gameover.atlas\"\n"
  "default_animation: \"Foxy Jump title\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 525.0
    y: 352.0
    z: 0.91
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  scale {
    x: 1.313
    y: 1.326
    z: 1.06
  }
}
embedded_components {
  id: "highscores"
  type: "sprite"
  data: "tile_set: \"/main/gameover/gameover.atlas\"\n"
  "default_animation: \"highscores title 3\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 525.0
    y: 108.0
    z: 0.9
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  scale {
    x: 0.617
    y: 0.617
    z: 0.617
  }
}
embedded_components {
  id: "playgame"
  type: "sprite"
  data: "tile_set: \"/main/gameover/gameover.atlas\"\n"
  "default_animation: \"play game title 2\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 525.0
    y: 210.0
    z: 0.91
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
