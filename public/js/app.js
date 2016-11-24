var GoDropBox = angular.module('GoDropBox', ['ngRoute']);


GoDropBox.config(function($routeProvider){
    $routeProvider.
           when('/',{
                templateUrl: 'pages/dragAndDrop.html',
                controller: 'homeController'
            })
            .when('/signin',{
                        templateUrl: 'pages/login.html',
                        controller: 'loginController'
                    })
            .when('/signup',{
                templateUrl: 'pages/register.html',
<<<<<<< HEAD
                controller: 'loginController'
            })
           
=======
                controller: 'homeController'
            })
            .when('/about',{
                templateUrl: 'pages/about.html',
                controller: 'homeController'
            });

>>>>>>> 5a4a3884f1a940d5c60968ee23026320d663293c
});

//main page controller
GoDropBox.controller('homeController', ['$scope', function($scope){
    console.log($scope);
}]);

GoDropBox.controller('loginController', ['$scope', function($scope){

}]);