/**
 * Consumer 消费
 */
var Consumer = {

    bindFancyBox: function() {

        $('[name="edit"]').each(function () {
            
            $(this).fancybox({
                padding: 12,
                minWidth: 500,
                minHeight: 456,
                width: '65%',
                height: '45%',
                autoSize: false,
                type: 'iframe',
                href: $(this).attr('data-link')
            });
        });

        $('button[name="add_consumer"]').each(function () {
            $(this).fancybox({
                padding: 12,
                minWidth: 500,
                minHeight: 456,
                width: '65%',
                height: '45%',
                autoSize: false,
                type: 'iframe',
                href: $(this).attr('data-link')
            });
        });
    },

    //定时更新状态
    status: function (url) {

        $.ajax({
            type : 'post',
            url : url,
            data : {'arr':''},
            dataType: "json",
            success : function(response) {
                if(response.code == 1) {
                    var values = response.data;
                    if (values.length == 0) {
                        return
                    }
                    for(var i = 0; i < values.length; i++) {
                        var statusElement = $('#consumer_'+values[i].consumer_id).find(".consumer_status");
                        if(values[i].status == 1) {
                            statusElement.html('<span class="label label-success">'+values[i].count+'</span>');
                        }else {
                            statusElement.html('<span class="label label-danger">'+values[i].count+'</span>');
                        }
                        var lastTimeElement = $('#consumer_'+values[i].consumer_id).find(".consumer_last_time");
                        lastTimeElement.html($.myTime.UnixToDate(values[i].last_time, true, 8))
                    }
                } else {
                    console.log(response);
                    if (response.redirect.url) {
                        console.log(response);
                        location.href = response.redirect.url;
                    }
                }
            },
            error : function(response) {
                console.log(response);
            }
        });
    }
};