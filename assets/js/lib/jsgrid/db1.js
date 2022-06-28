(function() {




    var db = {

        loadData: function(filter) {

            return $.grep(this.clients, function(client) {

                return (!filter.Name || client.Name.indexOf(filter.Name) > -1)
                    && (filter.CollegeName === undefined || client.Sex === filter.Sex)
                    && (!filter.Number || client.WorkId.indexOf(filter.WorkId) > -1);
            });
        },

        insertItem: function(insertingClient) {
            alert(111)
            this.clients.push(insertingClient);
        },

        updateItem: function(updatingClient) {
            $.ajax({
                url: '/changeRoom',
                type: 'POST',
                data:{
                    Id:updatingClient.Id,
                    Name:updatingClient.Name,
                    Number:updatingClient.Number,
                    CollegeName:updatingClient.CollegeName,


                },
                success: function(data){
                    if (data){
                        alert("修改成功")
                    }else {
                        alert("修改失败")
                    }
                }
            });
        },

        deleteItem: function(deletingClient) {
            $.ajax({
                url: '/delRoom',
                type: 'POST',
                data:{
                    Id:deletingClient.Id,

                },
                success: function(data){
                   if (data){
                       alert("删除成功")
                   }else {
                       alert("删除失败")
                   }
                }
            });
            var clientIndex = $.inArray(deletingClient, this.clients);
            this.clients.splice(clientIndex, 1);
        }

    };

    window.db = db;


    db.countries = college

    db.clients=rooms

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