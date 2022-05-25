#!/bin/sh
php artisan cache:clear
php artisan migrate --force
php artisan serve --host=0.0.0.0
