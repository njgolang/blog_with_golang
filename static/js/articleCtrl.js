// var more_article_path = "api/articles"
var more_article_path = "articles/20160101/test.md"
var app = angular.module('myApp', []);


app.factory('MDService', function($http, $q) {
	return {
		getData: function() {
			return $http.get("articles/20160101/test.md")
				.then(function(response) {
					return response;
				}, function(error) {
					console.log("get md file failed");
				})
		}
	}
})

app.filter('replaceImg', function() {
	return function(val) {
		var reg = new RegExp('', 'g');
		if(!val) val
	}
})

app.controller('getArticlesCtrl', ['$scope', '$http', '$q', function($scope, $http, $q) {
	var converter = new Showdown.converter();

	var defer = $q.defer();
	var promise = defer.promise;



    $http.get(more_article_path)
    	.success(function(response){
    		$scope.articles = response;

    		var htmlText = converter.makeHtml($scope.articles);
    		// element.html(htmlText);
    		
    	});

    $scope.abc = "abc";
    $scope.ab = "ab";

    // $http.get(more_article_path).success(function(response) {
    // 	alert($scope.abc);
    // 	$scope.articles = response;
    // })

   	$scope.loadDate = function() {
   		console.log("load data");
   	}

   	// $http.get(more_article_path).then(function(response) {
   	// 	$scope.articles = response;
   	// })
}]);

app.directive('markdown', function () {
    var converter = new Showdown.converter();
    return {
        restrict: 'E',
        link: function (scope, element, attr) {
        	var unwatch = scope.$watch('articles', function(v) {
        		if(v) {
        			unwatch();
        			var reg = /(!\[.*\])(\()(\S+)(\))/g;
        			console.log(v);
    				var data = v.replace(reg, '$1$2articles\/20160101\/$3$4')
    				console.log(data)
        			var htmlText = converter.makeHtml(data);
        			element.html(htmlText);
        		}
        	})
        	// scope.loadDate();
        	// function render() {
        	// 	MDService.getData()
        	// 	.then(function(data) {
        	// 		alert(data);
        	// 	});
        	// }

        	// render();

        	// alert(scope.articles);
      //   	$http.get(more_article_path)
		    // 	.success(function(response){
		    // 		scope.articles = response;
    		// })
        	// function renderMarkdown() {
         //    	var htmlText = converter.makeHtml(scope.articles);
         //    	element.html(htmlText);
        	// }
        	// scope.$watch(scope.articles, function() {
        	// 	renderMarkdown();
        	// });

        	// renderMarkdown();
        },
        // template: scope.articles



    };

});