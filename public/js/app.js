var GoDropBox = angular.module('GoDropBox', ['ngRoute']);


GoDropBox.config(function($routeProvider){
    $routeProvider.
        when('/',{
            templateUrl: 'pages/login.html',
            controller: 'homeController'
        });

});

//main page controller
GoDropBox.controller('homeController', ['$scope', function($scope){
    console.log($scope);
}]);

GoDropBox.controller('loginController', ['$scope', function($scope){

}]);