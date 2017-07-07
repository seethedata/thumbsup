var maxGas=4700000;

$().ready(function(){
	$("#enterForm").validate({
		rules: {
			applicationName: "required",
			approverType : {
				required: true,
				minlength: 1
			},
			environmentSize: "required"
		},
		messages:{
			applicationName: "Please enter a name.",
			approverType : "Please select at least 1 approver.",
			environmentSize: "Please select an environment size."
		},
		errorPlacement: function(error, element){
			error.insertBefore(element);
		}
	});
	$("#approveForm").validate({
		rules: {
			applicationName: "required",
			approverName : "required",
			approverType : "required",
		},
		messages:{
			applicationName: "Please select an application.",
			approverName : "Please enter a name.",
			approverType : "Please select an approver type.",
		},
		errorPlacement: function(error, element){
			error.insertBefore(element);
		}
	});
});

