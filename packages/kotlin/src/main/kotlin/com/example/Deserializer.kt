package com.example

import com.fasterxml.jackson.databind.ObjectMapper
import com.fasterxml.jackson.module.kotlin.readValue

class User(val name: String, val role: String)

fun main() {
    val mapper = ObjectMapper()
    val json = """{"name":"John","role":"user","@type":"com.example.User"}"""
    
    // Insecure deserialization vulnerability
    val user: User = mapper.readValue(json)
    println(user)
} 