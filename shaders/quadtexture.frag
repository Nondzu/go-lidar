#version 330 core

out vec4 FragColor;
in vec2 TexCoord;
in vec4 v_color;

uniform sampler2D texture1;


void main () {
  

  FragColor = texture(texture1, TexCoord); 

  FragColor *= v_color;
  // FragColor = vec4(1.0, 1.0, 0.0, 0.0);
  // vec4 ambientLight = vec4(1.0, 1.0,1.0,1.0);
  // vec4 ambientLight = vec4(0.3, 0.3, 0.3,1.0);
  // FragColor = ambientLight *  texture(texture1, TexCoord); 
  // FragColor = vec4(0.8, 0.30, 0.0, 1.0);



  // vec4 tex = texture(texture1, TexCoord);
  // FragColor = vec4(tex.x, tex.y, tex.z, 1.0);

}