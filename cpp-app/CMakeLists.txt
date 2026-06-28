cmake_minimum_required(VERSION 3.15)
project(CppSoapApp)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Declaración de dependencias del sistema
find_package(Crow REQUIRED)
find_package(gsoap REQUIRED)
find_package(cpr REQUIRED)

add_executable(v1_soap v1_soap.cpp)
add_executable(v2_traduccion v2_traduccion.cpp)
add_executable(v3_nativo v3_nativo.cpp)

# En Windows, vinculamos las librerías de red nativas del sistema operativo
if(WIN32)
    target_link_libraries(v1_soap PRIVATE ws2_32 wsock32)
    target_link_libraries(v2_traduccion PRIVATE ws2_32 wsock32)
    target_link_libraries(v3_nativo PRIVATE ws2_32 wsock32)
endif()

