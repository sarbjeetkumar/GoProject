var GoDropBox = angular.module('GoDropBox', ['ngRoute']);

GoDropBox.directive('jValidate', function() {
/*
This function is modified from - http://stackoverflow.com/questions/23707470/jquery-validation-is-not-working-with-templates-inside-ng-view

JQuery Validation will not work with Angular ng-view because validate is running before ng-view has rendered.
Directive is returned after ng-view renders
*/

    $.validator.addMethod( 'pattern', function( value, element) {
        return this.optional(element)
        || value.length >= 8 && /\d/.test(value) && /[!@#$&*]/i.test(value);
    }, "Password must be 8 characters long, contain one number and one special character" )

    $.validator.setDefaults({
    errorClass: 'help-block',
        highlight: function(element){
            $(element)
            .closest('.form-group')
            .addClass('has-error')
        },
          unhighlight: function(element){
                    $(element)
                    .closest('.form-group')
                    .removeClass('has-error')
                }
    })

  return {
    link: function(scope, element, attr) {

      element.validate({
    /*
    Validation is modified from tutorial - https://www.youtube.com/watch?v=xNSQ3i-BWMo
    */
            rules: {
                name:{
                    required: true
                },
                email: {
                    required:true,
                    email:true
                },
                username: {
                    required: true
                },
                password: {
                    required:true,
                    pattern:true
                },
                password2:{
                    required:true,
                    equalTo:"#password"
                },
                age: {
                    required: true
                },
                sex: {
                    required: true
                }
            },
            messages:{
                email:{
                    required: 'Please enter an email address.',
                    email: 'Please enter a valid email address.'
                },
                password2:{
                    required: 'Please enter password.',
                    equalTo: 'passwords do not match'
                }
            }
      });
    }
  }
});

GoDropBox.config(function($routeProvider){
    $routeProvider
           .when('/',{
                //templateUrl: 'pages/layout.tmpl',
                //templateUrl: 'pages/user.tmpl',
                //templateUrl: 'pages/upload.gtpl',
                templateUrl: 'pages/dragAndDrop.html',
                controller: 'homeController'
            })
           .when('/signin',{
                        templateUrl: 'pages/login.html',
                        controller: 'loginController'
            })
           .when('/signup',{
                templateUrl: 'pages/register.html',
                controller: 'loginController'

           })
           .when('/home',{
                        templateUrl: 'pages/dragAndDrop.html',
                        controller: 'homeController'
           })
          .when('/about',{
                templateUrl: 'pages/about.html',
                controller: 'aboutController'
            });

});

//main page controller
GoDropBox.controller('homeController', ['$scope', function($scope){
    console.log($scope);
}]);

GoDropBox.controller('loginController', ['$scope', function($scope){

}]);


GoDropBox.controller('aboutController', ['$scope', function($scope){

}]);

