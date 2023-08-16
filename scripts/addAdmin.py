import mysql.connector
import yaml
import bcrypt
import getpass

def readConfig():
    with open('config.yaml', 'r') as f:
        return yaml.safe_load(f)

def connectToDatabase():
    config = readConfig()
    try:
        connection = mysql.connector.connect(
            host=config['host'],
            user=config['dbUser'],
            password=config['password'],
            database=config['dbName'],
            port=config['port']
        )
        return connection
    
    except:
        return None


if __name__ == '__main__':
    connection = connectToDatabase()
    if connection:
        
        # Get admin credentials
        adminUserName = input("Enter an admin user name: ")
        adminPassword = getpass.getpass("Enter password: ")
        
        # Hash password
        saltrounds = 10
        hashedPassword = bcrypt.hashpw(adminPassword.encode("utf-8"), bcrypt.gensalt(rounds = saltrounds))
        
        data = {
            "username": adminUserName,
            "hash": hashedPassword
        }
        
        # Adding a new admin to db.
        try: 
            cursor = connection.cursor()
            cursor.execute('''INSERT INTO users (username, admin, hash, requestAdmin) 
                        VALUES (%(username)s, 1, %(hash)s, 0)''', data)
            connection.commit()
            print(f"\033[1mSuccessfully registered '{adminUserName}' as an admin :) \033[0m")
            
        except mysql.connector.Error as err:
            print(f"\033[1mError occured while adding '{adminUserName}' to database: {err} \033[0m")
        
    else:
        print("\033[1mError connecting to database. Please check the credentials.\033[0m")