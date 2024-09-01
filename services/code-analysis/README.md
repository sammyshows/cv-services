# Code Analysis Service

## Purpose
This service performs deep code analysis, transforms codebases into ASTs, and populates the graph database with code structure and metadata. It supports various programming languages and provides the foundational data needed for the Documentation Service to generate meaningful insights and documentation.

## Features
- Source code parsing to AST
- AST-based code analysis
- Graph database integration for code structure
- Support for multiple programming languages

## Technology Stack
- Node.js for processing and logic
- Babel (for JavaScript), RubyParser (for Ruby), and others for parsing
- Neo4j for storing and querying code data
- Docker for consistent deployment environments
