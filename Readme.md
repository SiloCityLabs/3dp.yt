QR Code URL Redirector
======================

This project provides a simple URL redirector that automatically redirects users from a QR code URL to the appropriate page on the filameter.com website based on specific URL path formats.

### Features

*   Redirects users based on the URL path.
    
*   Supports various path formats, including:
    
    *   **SpoolSense Display QR**: Redirects with the format /s/{id}/{used}.
        
    *   **Filameter QR Labels**: Redirects with the format /fq/{id}.
        
    *   **Custom QR Labels**: Redirects with the format /custom/{id}.
        
*   Displays an error message if the URL format is invalid.
    

### Usage

When a user scans the QR code, the URL will be redirected based on the following rules:

1.  **SpoolSense Display QR**: If the path starts with /s/, the URL will be redirected to https://filameter.com/ssi?id={id}&used={used}.
    
2.  **Filameter QR Labels**: If the path starts with /fq/, the URL will be redirected to https://filameter.com/s?id={id}.
    
3.  **Custom QR Labels**: If the path starts with /custom/, the URL will be redirected to https://filameter.com/s?id={id}.
    
4.  **Invalid URL Format**: If the path does not match any of the above patterns, an error message is displayed.
    

### Installation

To use this redirector on your own site, copy the HTML code and integrate it into your website's page. Ensure that the page is accessible via the QR code URLs.

### Example

For a URL like:
```
https://yoursite.com/s/12345/true
```

The user will be redirected to:
```
https://filameter.com/ssi?id=12345&used=true
```

For a URL like:
```
https://yoursite.com/fq/67890
```

The user will be redirected to:
```
https://filameter.com/s?id=67890
```

### Notes

*   The URL must include specific path segments as mentioned above (e.g., /s/{id}/{used}, /fq/{id}, or /custom/{id}).
    
*   If the URL doesn't match the expected format, a simple error message will be displayed.
    
*   The script automatically handles the redirection or shows an error message.
    

### License

This project is licensed under the MIT License. See the LICENSE file for more information.