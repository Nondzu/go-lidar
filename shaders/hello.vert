#version 330 core
layout (location = 0) in vec3 aPos;
layout (location = 1) in vec2 aTexCoord;

out vec2 TexCoord;
out vec4 v_color; 

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;
uniform vec4 color;

void main() 
{
    // v_color = vec4(0.0, 1.0, 0.0, 1.0);
    // v_color = vec4(0.5, 0.5, 0.0, 1.0);
    v_color = color;
    // gl_Position = vec4(aPos.x,aPos.y,aPos.z,1.0f);
    gl_Position = projection * view * model * vec4(aPos, 1.0f);
    TexCoord = vec2(aTexCoord.x, 1.0f - aTexCoord.y);
}
