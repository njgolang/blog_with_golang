var app = angular.module('myApp', []);


app.controller('pageCtrl', ['$scope', '$http', '$window', function($scope, $http, $window) {
	var path = $window.location.pathname;
    var mdPath = path.replace('page', 'articles');

    var arr = mdPath.split('/');

    $http.get(mdPath)
    	.success(function(response){
    		$scope.articles = response;
            $scope.mdTimeDir = arr[arr.length -2]

    	});
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
        			// console.log(v);
    				var data = v.replace(reg, '$1$2\/articles\/' + scope.mdTimeDir + '\/$3$4')
    				console.log(data)
        			var htmlText = converter.makeHtml(data);
        			element.html(htmlText);
        		}
        	})
        },
    };
});