from flask import Flask, request, jsonify
import psycopg2
import os

app = Flask(__name__)

def get_db_connection():
    return psycopg2.connect(
        dbname=os.environ['DB_NAME'],
        user=os.environ['DB_USER'],
        password=os.environ['DB_PASS'],
        host=os.environ['DB_HOST'],
        port=os.environ['DB_PORT']
    )

@app.route('/write', methods=['POST'])
def write_data():
    data = request.json.get('data')
    conn = get_db_connection()
    cur = conn.cursor()
    cur.execute("INSERT INTO demo_table (data) VALUES (%s)", (data,))
    conn.commit()
    cur.close()
    conn.close()
    return jsonify({'message': 'Data written!'})

@app.route('/read')
def read_data():
    conn = get_db_connection()
    cur = conn.cursor()
    cur.execute("SELECT data FROM demo_table")
    rows = cur.fetchall()
    cur.close()
    conn.close()
    return jsonify([row[0] for row in rows])

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
