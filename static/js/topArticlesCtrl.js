var topArticles = "api/articles"

var app = angular.module('myApp', []);

app.controller('topArticlesCtrl', ['$scope', '$http', '$window', function($scope, $http, $window) {
    $http.get(topArticles)
    	.success(function(response){
    		$scope.articles = response;
    	});

    $scope.go = function(Path) {
    	var pagePath = Path.replace('articles', 'page')
    	$window.location.href = pagePath;
    }
}]);