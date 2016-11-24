var GoDropBox = angular.module('GoDropBox', ['ngRoute']);


GoDropBox.config(function($routeProvider){
    $routeProvider.
           when('/',{
                templateUrl: 'pages/dragAndDrop.html',
                controller: 'homeController'
            })
            .when('/signin',{
                        templateUrl: 'pages/login.html',
                        controller: 'homeController'
                    })
            .when('/signup',{
                templateUrl: 'pages/register.html',
                controller: 'homeController'
            })
            .when('/about',{
                templateUrl: 'pages/about.html',
                controller: 'homeController'
            });

});

//main page controller
GoDropBox.controller('homeController', ['$scope', function($scope){
    console.log($scope);
}]);

GoDropBox.controller('loginController', ['$scope', function($scope){

}]);