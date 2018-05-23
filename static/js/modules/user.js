/**
 * user 用户
 */
var User = {

    bindFancyBox: function() {

        $('[name="edit"]').each(function () {
            $(this).fancybox({
                minWidth: 450,
                minHeight: 200,
                padding: 12,
                width: '65%',
                height: '52%',
                autoSize: false,
                type: 'iframe',
                href: $(this).attr('data-link')
            });
        });

        $('[name="node"]').each(function () {
            $(this).fancybox({
                minWidth: 650,
                minHeight: 400,
                padding: 12,
                width: '70%',
                height: '52%',
                autoSize: false,
                type: 'iframe',
                href: $(this).attr('data-link')
            });
        });
    },

    selectNode: function (element) {
        $(element).bind('click', function () {
            var checked = $(this).is(':checked');
            console.log(element);
            $('input[type="checkbox"][name="node_id"]').each(function() {
                this.checked = checked
            });
        });
    }
};