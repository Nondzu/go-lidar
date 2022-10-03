#version 330 core

out vec4 FragColor;

in vec2 TexCoord;
in vec3 Normal;
in vec3 FragPos;

uniform sampler2D texture1;

uniform vec3 lightPos;
uniform vec3 lightColor;  // color/brightness
uniform vec3 ambientColor;  // color/brightness

void main () {

  //ambient  
  // vec3 ambientLight = vec3(0.1, 0.1, 0.1);

  //diffuse
  vec3 lightDir = normalize(lightPos - FragPos);
  float diff = max(dot(Normal, lightDir), 0.0);
  vec3 diffuse = diff * lightColor;
    
  FragColor = vec4(ambientColor + diffuse, 1.0) * texture(texture1, TexCoord);
}