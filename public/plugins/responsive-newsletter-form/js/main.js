jQuery(document).ready(function($){
	var messages = $('div[data-type="message"]');
	//check if user updates the email field
	$('.cd-form .cd-text').keyup(function(event){	
		//check if user has pressed the enter button (event.which == 13)
		if(event.which!= 13) {
			//if not..
			//hide messages and loading bar 
			messages.removeClass('slide-in is-visible');
			$('.cd-form').removeClass('is-submitted').find('.cd-loading').off('webkitTransitionEnd otransitionend oTransitionEnd msTransitionEnd transitionend');
		}

		var TextInput = $(this),
			insertedText = TextInput.val();
			
	    //check if user has inserted a "@" and a dot
	    if (insertedText == '') {
	    	//if he hasn't..
	    	//hide the submit button
	    	$('.cd-form').removeClass('is-active').find('.cd-loading').off('webkitTransitionEnd otransitionend oTransitionEnd msTransitionEnd transitionend');
	    } else {
	    	//if he has..
	    	//show the submit button
	    	$('.cd-form').addClass('is-active');
	    }
	});

	//backspace doesn't fire the keyup event in android mobile
	//so we check if the email input is focused to hide messages and loading bar 
	$('.cd-form .cd-text').on('focus', function(){
		messages.removeClass('slide-in is-visible');
		$('.cd-form').removeClass('is-submitted').find('.cd-loading').off('webkitTransitionEnd otransitionend oTransitionEnd msTransitionEnd transitionend');
	});	

	//you should replace this part with your ajax function
	$('.cd-submit').on('click', function(event){
		var url = $('.cd-form').attr("action"),
			content = $("input[name='content']").val(),
			title = $("input[name='title']").val(),
			// subtitle = $("input[name='subtitle']").val(),
			data = {
				content: content,
				title: title
				// subtitle: subtitle
			};
		if($('.cd-form').hasClass('is-active')) {
			event.preventDefault();
			//show the loading bar and the corrisponding message
			$('.cd-form').addClass('is-submitted').find('.cd-loading').one('webkitTransitionEnd otransitionend oTransitionEnd msTransitionEnd transitionend', function(){
				$.ajax({
  					type: "POST",
  					url: url,
  					data: data,
  					success: function (result) {
  						showMessage(result.status.code, result.status.msg);
  						//if transitions are not supported - show messages
						if($('html').hasClass('no-csstransitions')) {
							showMessage(result.status.code, result.status.msg);
						}
  					},
 	 				dataType: "json"
				});
			});
		}
	});

	function showMessage(code, msg) {
		if(code == 1001) {
			$('.cd-response-success').addClass('slide-in');
		} else {
			if (msg != '') {
				$('.cd-response-error').html(msg)
			}
			$('.cd-response-error').addClass('is-visible');
		} 
	}

	//placeholder fallback (i.e. IE9)
	//credits http://www.hagenburger.net/BLOG/HTML5-Input-Placeholder-Fix-With-jQuery.html
	if(!Modernizr.input.placeholder){
		$('[placeholder]').focus(function() {
			var input = $(this);
			if (input.val() == input.attr('placeholder')) {
				input.val('');
		  	}
		}).blur(function() {
		 	var input = $(this);
		  	if (input.val() == '' || input.val() == input.attr('placeholder')) {
				input.val(input.attr('placeholder'));
		  	}
		}).blur();
		$('[placeholder]').parents('form').submit(function() {
		  	$(this).find('[placeholder]').each(function() {
				var input = $(this);
				if (input.val() == input.attr('placeholder')) {
			 		input.val('');
				}
		  	})
		});
	}
});