# Third-pary stuff
This directory contains ("vendors") abbreviated copies of the following repositories:

* googleapis/google - https://github.com/googleapis/googleapis/ 7cde5d0df08801e00fc45df1546942aa2692d5c3 (LICENSE, google/api, google/rpc)
* grpc-gateway - https://github.com/grpc-ecosystem/grpc-gateway 1dd92c92ad0d0f78c750d0e41f8ff91edad99cc2 (LICENSE.txt, protoc-gen-openapiv2/options)
* OpenAPI - https://github.com/swagger-api/swagger-ui 07a0416ff664583ff9f481cae7dace226c9f61ec (LICENSE, dist/)

The `third_party/OpenAPI` directory contains HTML, Javascript,
and CSS assets that dynamically generate Swagger documentation from a
Swagger-compliant API definition file. That file is auto-generated.
The static assets are copied from [this dist folder](https://github.com/swagger-api/swagger-ui/tree/master/dist)
of the OpenAPI-UI project. After copying, [index.html](./OpenAPI/index.html)
is edited to load the generated swagger file from the local server instead of the default petstore
and the edited version is stored [here](./openapiv2/v1/index.html).

I've also copied the OpenAPI files in the `openapiv2/v1` directory.

See the respective LICENSE files for each project for the applicable license terms.
