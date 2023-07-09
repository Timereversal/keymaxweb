    

    - CarKey information
        - Car Brand
        - Car Model
        - Car Year
        - 

    

    - Register Car key information / sales o not.
        - ID
        - Customer Information [Name, Cellphone]
        - Vehicule Information [Brand, Model, Year,VIN, Placa]
        - Key Information [Blade, Transponder, FCC-OEM PartNumber, Pictures]
        - Service information [Type of service, like chip programming, key shell change, etc.]
        - Products [Blade, Transponder, Remote,PortaChip]
        - Additional Comments. [Text,Pictures,Tools]
        - Date 
        - Sale Information [sale, monto]

        * Add  /v1/customerkey

        Models:
            Car:
                - Brand
                - Model
                * Add       /v1/car/add
                * Edit      /v1/car/edit
                * Delete.   /v1/car/delete

            Service:
                - type [Change key shell, Chip programming, Reprogramming]
                * add
                * delete
                * update 

    -Pedidos 
    -Inventory

        - Products:

            - Key Fob Remotes
            - Key Shells
            - Blades
                - Key Fob Model
                    * id
                    * sku
                    * type string # wire, superremote,NXP,NB,ZB,XM28,XM38Toyota
                    * brand string # KD,Xhorse,JMD,Lonsdor,Aftermarket,OEM
                    * seller
                    * quantity
                    * model
                    * frequency
                    * transponder
                    * blade

            - Transponders
            

    - Create Table
        - User
        - 
    - EndPoints.
        * POST  /carkeyinfo/post/
        * GET   /carkeyinfo/list (support pagination)
        * GET   /carkeyinfo/byID=123  (show all information related to one )

