/**
 * 首页
 * Copyright (c) 2017 phachon@163.com
 */

var Main = {
	/**
	 * 获取服务器状态
     * @param url
	 * @constructor
	 */
	GetServerStatus: function (url) {
		$.ajax({
			type : 'post',
			url : url,
			data : {'arr':''},
			dataType: "json",
			success : function(response) {
				if(response.code == 0) {
					console.log(response.message);
					return false
				}
				var cpu = response.data.cpu_used_percent;
				var memory = response.data.memory_used_percent;
				var disk = response.data.disk_used_percent;
				// cpu
				$(".cpu_text").each(function () {
					$(this).text(cpu+"%")
				});
				$("#cpu_progress").attr("aria-valuenow", cpu);
				$("#cpu_progress").attr('style', 'min-width: 2em; width: '+cpu+'%');

				// memory
				$(".memory_text").each(function () {
					$(this).text(memory+"%")
				});
				$("#memory_progress").attr("aria-valuenow", memory);
				$("#memory_progress").attr('style', 'min-width: 2em; width: '+memory+'%');

				// disk
				$(".disk_text").each(function () {
					$(this).text(disk+"%")
				});
				$("#disk_progress").attr("aria-valuenow", disk);
				$("#disk_progress").attr('style', 'min-width: 2em; width: '+disk+'%');

			},
			error : function(response) {
				console.log(response.message)
			}
		});
	}
};