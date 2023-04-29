$(function () {
	app.init()
})

var app = {
	init: function () {
		this.getCaptcha()
		this.changeCaptcha()
	},
	getCaptcha: function () {
		$.get('/admin/captcha?t=' + Math.random(), function (response) {
			// console.log(response)
			$('#captcha_id').val(response.captcha_id)
			$('#captcha_image').attr('src', response.captcha_image)
		})
	},
	changeCaptcha: function () {
		var that = this
		$('#captcha_image').click(function () {
			that.getCaptcha()
		})
	}
}
