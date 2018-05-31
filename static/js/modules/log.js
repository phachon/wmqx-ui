var Log = {
	
	bindFancyBox: function() {
		$('button[name="log_download"]').each(function () {
			$(this).fancybox({
				minWidth: 600,
				minHeight: 420,
				width: '80%',
				height: '45%',
				autoSize: false,
				type: 'iframe',
				href: $(this).attr('data-link')
			});
		});
        $('a[name="info"]').each(function () {
            $(this).fancybox({
				padding: 12,
                minWidth: 400,
                minHeight: 310,
                width: '60%',
                height: '40%',
                autoSize: false,
                type: 'iframe',
                href: $(this).attr('data-link')
            });
        });
	}
};
