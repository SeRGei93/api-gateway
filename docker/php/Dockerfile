FROM php:8.4-fpm-alpine

WORKDIR /var/www

RUN apk add --no-cache bash curl git unzip

# Устанавливаем Composer
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/bin --filename=composer

EXPOSE 9000

CMD ["php-fpm"]