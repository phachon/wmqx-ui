/**
 * node 节点
 */
var Node = {

    bindFancyBox: function() {

        $('[name="edit"]').each(function () {
            $(this).fancybox({
                minWidth: 500,
                minHeight: 370,
                padding: 12,
                width: '65%',
                height: '48%',
                autoSize: false,
                type: 'iframe',
                href: $(this).attr('data-link')
            });
        });

        $('[name="message"]').each(function () {
            $(this).fancybox({
                minWidth: 500,
                minHeight: 370,
                padding: 12,
                width: '65%',
                height: '48%',
                autoSize: false,
                type: 'iframe',
                href: $(this).attr('data-link')
            });
        });
    },

    GetNodeStatus: function (url) {

        $("td[data-name='manager_ri']").each(function () {
            var nodeId = $(this).parent().attr("data-row");
            request(url, nodeId);
        });

        function request(url, nodeId) {
            $.ajax({
                type : 'get',
                url : url,
                async: true,
                data : {'node_id': nodeId},
                dataType: "json",
                success : function(response) {
                    if(response.code == 1) {
                        var version = response.data.version;
                        $("#status_"+nodeId).html("<span class='label label-success'>"+version+"</span>")
                    }else {
                        $("#status_"+nodeId).html("<span class='label label-danger'>Error</span>")
                    }
                },
                error : function(response) {
                    console.log(response.message)
                }
            });
        }
    }
};