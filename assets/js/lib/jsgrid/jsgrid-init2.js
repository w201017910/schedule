$(function() {

    $("#jsGrid").jsGrid({
        height: "100%",
        width: "100%",
        filtering: false,
        editing: true,
        inserting: false,
        sorting: true,
        paging: true,
        autoload: true,
        pageSize: 15,
        pageButtonCount: 5,
        deleteConfirm: "你确定要删除？",
        controller: db,
        fields: [
            { name: "Name", type: "text", width: 150,title: "班级名" },
            { name: "Number", type: "text", width: 50,title: "班级人数" },
            { name: "college", type: "select", items: db.countries, valueField: "college", textField: "college",title: "所属学院" },
            { name: "room", type: "select", items: db.countrie, valueField: "DefaultRoom", textField: "room",title: "默认班级" },


            { type: "control" }
        ]
    });

});