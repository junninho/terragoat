package com.example;

import org.springframework.web.bind.annotation.*;
import java.io.File;
import java.nio.file.Files;

@RestController
public class FileController {
    
    @GetMapping("/file")
    public String getFile(@RequestParam String filename) {
        // Path traversal vulnerability
        File file = new File("/var/www/files/" + filename);
        try {
            return new String(Files.readAllBytes(file.toPath()));
        } catch (Exception e) {
            return "Error reading file";
        }
    }
} 