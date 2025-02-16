package com.example.demo.service;

import com.example.demo.model.User;
import org.springframework.stereotype.Service;
import java.util.ArrayList;
import java.util.List;

@Service  // Marks this class as a service component
public class UserService {
    private final List<User> users = new ArrayList<>();

    // Create User
    public User addUser(User user) {
        users.add(user);
        return user;
    }

    // Get All Users
    public List<User> getAllUsers() {
        return users;
    }

    // Get User by ID
    public User getUserById(Long id) {
        return users.stream().filter(user -> user.getId().equals(id)).findFirst().orElse(null);
    }

    // Update User
    public User updateUser(Long id, User updatedUser) {
        for (User user : users) {
            if (user.getId().equals(id)) {
                user.setName(updatedUser.getName());
                user.setEmail(updatedUser.getEmail());
                return user;
            }
        }
        return null;
    }

    // Delete User
    public boolean deleteUser(Long id) {
        return users.removeIf(user -> user.getId().equals(id));
    }
}
