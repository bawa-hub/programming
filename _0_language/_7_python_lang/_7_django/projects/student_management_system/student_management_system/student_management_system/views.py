from django.shortcuts import render,redirect,HttpResponse
from app.EmailBackend import EmailBackend
from django.contrib.auth import authenticate,login as auth_login,logout
from django.contrib import messages
from django.contrib.auth.decorators import login_required


def base(request):
    return render(request, 'base.html')

def login(request):
    return render(request, 'login.html')

def doLogin(request):
    if request.method == 'POST':
        user = EmailBackend.authenticate(request, username=request.POST.get('email'),
                                         password=request.POST.get('password'))
        if(user is not None):
            auth_login(request, user)
            user_type=user.user_type
            if user_type == '1':
                return redirect('hod_home')
            elif user_type == '2':
                return HttpResponse("This is Staff panel")
            elif user_type == '3':
                return HttpResponse("This is Student panel")
            else:
                messages.error(request, "Email and Password are invalid.")
                return redirect('login')
        else:
            messages.error(request, "Email and Password are invalid.")
            return redirect('login')    

def doLogout(request):
    logout(request)
    return redirect('login')        
        