# URL shortener server

#### Dependencies
* github.com/julienschmidt/httprouter **v1.3.0**

#### Handlers
* RedirectByShortenLink **GET**
    
    Return codes
    * StatusFound **302**
    * StatusNotFound **404**
    
* GenerateShortenLink **POST**

    Return codes
    * StatusCreated **201**
    * StatusBadRequest **400**