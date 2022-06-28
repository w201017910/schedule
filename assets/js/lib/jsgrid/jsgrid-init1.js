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
            { name: "Name", type: "text", width: 150,title: "教室名" },
            { name: "Number", type: "text", width: 50,title: "容纳量" },
            { name: "CollegeName", type: "select", items: db.countries, valueField: "Name", textField: "Name",title: "所属学院" },



            { type: "control" }
        ]
    });

});