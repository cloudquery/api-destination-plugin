from http.server import BaseHTTPRequestHandler, HTTPServer
import logging
import json

class RequestHandler(BaseHTTPRequestHandler):
    def _set_response(self):
        self.send_response(200)
        self.send_header('Content-type', 'application/json')
        self.end_headers()

    def do_POST(self):
        if self.path not in ['/append', '/migrate']:
            self.send_error(404, "Endpoint not found")
            return
        
        content_length = int(self.headers['Content-Length'])  # Get the size of the data
        post_data = self.rfile.read(content_length)  # Get the data

        try:
            json_data = json.loads(post_data)
        except json.JSONDecodeError:
            self.send_error(400, "Invalid JSON")
            return

        # Print headers and body
        logging.info("Headers: %s", str(self.headers))
        logging.info("Body: %s", json.dumps(json_data, indent=2))

        # Set response
        self._set_response()
        response = {
            "message": "Request received",
            "path": self.path
        }
        self.wfile.write(json.dumps(response).encode('utf-8'))

def run(server_class=HTTPServer, handler_class=RequestHandler, port=8080):
    logging.basicConfig(level=logging.INFO)
    server_address = ('', port)
    httpd = server_class(server_address, handler_class)
    logging.info('Starting httpd...\n')
    httpd.serve_forever()

if __name__ == '__main__':
    from sys import argv

    if len(argv) == 2:
        run(port=int(argv[1]))
    else:
        run()
