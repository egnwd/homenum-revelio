'use strict';

var gulp = require('gulp');
var sass = require('gulp-sass');
var typings = require('gulp-typings');
var gutil = require('gulp-util');
var uglify = require('gulp-uglify');
var browserify = require("browserify");
var source = require('vinyl-source-stream');
var tsify = require("tsify");

var targetDir = __dirname + '/dist';

gulp.task('typings', function(){
    return gulp.src('./typings.json')
        .pipe(typings());
});

var tsSource = browserify({
    basedir: '.',
    debug: true,
    entries: ['src/index.tsx'],
    cache: {},
    packageCache: {}
}).plugin(tsify);

gulp.task('scripts', ['typings'], function() {
    return tsSource.bundle()
    .pipe(source('./js/bundle.js'))
    .pipe(gulp.dest(targetDir));
});

gulp.task('styles', function () {
    return gulp.src('./style/**/*.scss')
        .pipe(sass().on('error', sass.logError))
        .pipe(gulp.dest(targetDir + '/css'));
});

gulp.task('html', function () {
    return gulp.src('*.html')
        .pipe(gulp.dest(targetDir));
});

gulp.task('all', ['scripts', 'styles', 'html']);
gulp.task('default', ['all']);
