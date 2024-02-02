# Official Nginx image
FROM nginx:latest

# Remove default nginx welcome page exposed on port 80
RUN mkdir -p /etc/nginx/sites-enabled/ \
    && ln -s /etc/nginx/sites-available/symfony.conf /etc/nginx/sites-enabled/symfony.conf \
    && rm /etc/nginx/conf.d/default.conf