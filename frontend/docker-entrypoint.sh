# docker-entrypoint.sh
#!/bin/bash

# Wait for database to be ready
echo "Waiting for database..."
while ! nc -z db 3306; do
  sleep 1
done
echo "Database is ready!"

# Run migrations if needed
if [ "$APP_ENV" = "production" ]; then
    echo "Running migrations..."
    php artisan migrate --force
fi

# Clear and cache configuration
php artisan config:clear
php artisan config:cache
php artisan route:clear
php artisan route:cache
php artisan view:clear
php artisan view:cache

# Set proper permissions
chown -R www-data:www-data /var/www
chmod -R 755 /var/www
chmod -R 775 /var/www/storage
chmod -R 775 /var/www/bootstrap/cache

# Start supervisor
exec "$@"