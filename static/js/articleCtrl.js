var more_article_path = "api/articles"
var app = angular.module('mainApp', []);

app.controller('getArticlesCtrl', ['$scope', '$http', function($scope, $http) {
    $http.get(more_article_path)
        .success(function(response) {
            $scope.articles = response;
    });
}]);
