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
    }
};