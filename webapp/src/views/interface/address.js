$(function () {
    var addressGrid;
    var ifaceList;

    ifaceList = $('#ifaceList').datalist({
        method: 'GET',
        url: '/api/interfaces',
        valueField: 'name',
        textField: 'name',
        lines: true,
        onLoadSuccess: function () {
            ifaceList.datalist('selectRow', 0);
        },
        onSelect: function (index, iface) {
            addressGrid.datagrid({
                url: '/api/interfaces/' + iface.type + '/' + iface.name + '/address'
            });
        }
    });

    var toollbar = [{
        text: '<i class="fa fa-plus-square" aria-hidden="true"></i> 添加',
        handler: function () { alert('add') }
    }, {
        text: '<i class="fa fa-trash" aria-hidden="true"></i> 删除',
        handler: function () { alert('delete') }
    }, {
        text: '<i class="fa fa-refresh" aria-hidden="true"></i> 刷新',
        handler: function () {
            grid.datagrid('reload');
        }
    }];
    var columns = [[
        { field: 'checkbox', checkbox: true },
        { field: 'local', title: 'IP', width: 200, align: 'left', sortable: true },
        { field: 'prefixlen', title: '掩码', width: 100, align: 'left' },
        {
            field: 'dynamic', title: '动态地址', width: 100, align: 'center',
            formatter: function (dynamic, row, index) {
                return dynamic ? '是' : '否';
            }
        }
    ]];

    addressGrid = $('#addressGrid').datagrid({
        toolbar: toollbar,
        columns: columns,
        fitColumns: true,
        method: 'GET',
        loadMsg: '处理中，请等待...',
        remoteSort: false,
        striped: true,
        pagination: true,
        pageSize: 50,
        pageList: [20, 50, 100, 500]
    });
});