(function() {
    // $.ajax({
    //     url: '/allTeacher',
    //     type: 'POST',
    //     data:{
    //     },
    //     success: function(data){
    //         db.clients=data
    //     }
    // });




    var db = {

        loadData: function(filter) {

            return $.grep(this.clients, function(client) {

                return (!filter.Name || client.Name.indexOf(filter.Name) > -1)
                    && (filter.Sex === undefined || client.Sex === filter.Sex)
                    && (!filter.WorkId || client.WorkId.indexOf(filter.WorkId) > -1)
                    && (!filter.Phone || client.Phone === filter.Phone)
                    && (filter.Email === undefined || client.Email === filter.Email);
            });
        },

        insertItem: function(insertingClient) {
            this.clients.push(insertingClient);
        },

        updateItem: function(updatingClient) {
            $.ajax({
                url: '/changeTeacher',
                type: 'POST',
                data:{
                    Id:updatingClient.Id,
                    Name:updatingClient.Name,
                    Sex:updatingClient.Sex,
                    Phone:updatingClient.Phone,
                    Email:updatingClient.Email,
                    WorkId:updatingClient.WorkId,

                },
                success: function(data){
                    alert(data)
                }
            });
        },

        deleteItem: function(deletingClient) {
            var clientIndex = $.inArray(deletingClient, this.clients);
            this.clients.splice(clientIndex, 1);
        }

    };

    window.db = db;


    db.countries = [
        { Name: "女"},
        { Name: "男"},

    ];

    db.clients=teachers

    db.users = [
        {
            "ID": "x",
            "Account": "A758A693-0302-03D1-AE53-EEFE22855556",
            "Name": "Carson Kelley",
            "RegisterDate": "2002-04-20T22:55:52-07:00"
        },
        {
            "Account": "D89FF524-1233-0CE7-C9E1-56EFF017A321",
            "Name": "Prescott Griffin",
            "RegisterDate": "2011-02-22T05:59:55-08:00"
        },

     ];

}());