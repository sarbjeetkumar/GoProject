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

GoDropBox.config(function($routeProvider, $locationProvider){
    $routeProvider
           .when('/',{
                templateUrl: 'pages/landing.html',
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
        /* .when('/about',{
                templateUrl: 'pages/about.html',
                controller: 'aboutController'
            })*/
            .when('/fileupload',{
                templateUrl: 'pages/dragAndDrop.html',
                controller: 'dragAndDropController'
            })
           .when('/home',{
                templateUrl: 'pages/landing.html',
                controller: 'homeController'
           });
});

//main page controller
GoDropBox.controller('homeController', ['$scope', function($scope){

    console.log($scope);
}]);

<<<<<<< HEAD


GoDropBox.controller('loginController', ['$scope', function($scope){
=======
GoDropBox.controller('loginController', ['$scope', '$location', function($scope, $location){
>>>>>>> 0705d4c07911c3a8ca7d6f28b6af508abe3a414b
    console.log($scope);
    //http://stackoverflow.com/questions/14201753/angular-js-how-when-to-use-ng-click-to-call-a-route
    $scope.go = function ( path ) {
        console.log(path);
     $location.path( path );
    };
}]);


GoDropBox.controller('aboutController', ['$scope', function($scope){

}]);

GoDropBox.controller('dragAndDropController', ['$scope', '$location', function($scope, $location){

}]);

