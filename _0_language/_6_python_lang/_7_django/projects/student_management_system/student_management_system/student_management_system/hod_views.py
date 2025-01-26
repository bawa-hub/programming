from django.shortcuts import render,redirect
from django.contrib.auth.decorators import login_required

@login_required(login_url='/')
def home(request):
    # print("isauth", request.user.is_authenticated)  # Debugging print
    return render(request, 'Hod/home.html')