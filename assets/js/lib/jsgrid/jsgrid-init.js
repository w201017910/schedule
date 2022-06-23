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
                    { name: "Name", type: "text", width: 150,title: "姓名" },
                    { name: "Sex", type: "select", items: db.countries, valueField: "Name", textField: "Name",title: "性别" },
                    { name: "WorkId", type: "text", width: 50,title: "学工号" },
                    { name: "Phone", type: "text", width: 200 ,title: "电话"},
                    { name: "Email", type: "text", width: 200 ,title: "电子邮件" },

                    { type: "control" }
                ]
            });

        });