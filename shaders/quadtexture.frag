#version 330 core

out vec4 FragColor;
in vec2 TexCoord;

uniform sampler2D texture1;


void main () {
  

  FragColor = texture(texture1, TexCoord); 
  // vec4 tex = texture(texture1, TexCoord);
  // FragColor = vec4(tex.x, tex.y, tex.z, 1.0);

}