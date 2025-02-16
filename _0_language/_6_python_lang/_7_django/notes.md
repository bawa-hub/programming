# create virtual environment

python3 -m venv myvenv

# activate venv

source /path/to/myvenv/bin/activate

# deactivate venv
deactivate

# install dependencies

pip3 install -r requirements.txt

# start django project

pip install django
django-admin startproject <projectname>

# make apps

cd project
python manage.py startapp <appname>c

# start server

python manage.py runserver

# make and run migration

python manage.py makemigrations
python manage.py migrate

# run django shell

python manage.py shell

# django admin panel

python manage.py createsuperuser

# install or make requirements.txt file

-- activate venv
-- pip install -r requirements.txt for installing
-- pip freeze > requirements.txt for making 
