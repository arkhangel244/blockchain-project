from flask import Flask, request, jsonify
from enum import Enum



class BankName(Enum):
    CHASE = "Chase Bank"
    SBI = "State Bank of India"
    BANK_OF_AMERICA = "Bank of America"
    CITIBANK = "Citibank"
    HSBC = "HSBC Bank"
    BARCLAYS = "Barclays Bank"
    JP_MORGAN = "J.P. Morgan Chase Bank"
    STANDARD_CHARTERED = "Standard Chartered Bank"

class Country(Enum):
    USA = "United States"
    UK = "United Kingdom"
    CANADA = "Canada"
    GERMANY = "Germany"
    FRANCE = "France"
    AUSTRALIA = "Australia"
    JAPAN = "Japan"
    INDIA = "India"
    CHINA = "China"

class Currency(Enum):
    USD = "US Dollar"
    GBP = "British Pound"
    CAD = "Canadian Dollar"
    EUR = "Euro"
    AUD = "Australian Dollar"
    JPY = "Japanese Yen"
    INR = "Indian Rupee"
    CNY = "Chinese Yuan"


class Customer:
    def __init__(self, username, password):
        self.username = username
        self.password = password
        self.account_details = {}

    def __str__(self):
        return f"Username: {self.username}, Account Details: {self.account_details}"

    def register_to_bank(self, bank, account_password, bank_name, amount):
        if account_password != self.password:
            return "Invalid account password"

        if bank_name not in bank_servers:
            return "Bank not found"

        bank_server = bank_servers[bank_name]

        if self.username in bank_server.users:
            return "User already registered to the bank"

        self.account_details['balance'] = amount
        bank_server.users[self.username] = self

        return "Successfully registered to the bank and amount added to the account"
    
    
    
class Bank:
    def __init__(self, name, country, preferred_currency, operates_in):
        self.name = name
        self.country = country
        self.preferred_currency = preferred_currency
        self.operates_in = operates_in
        self.currency_preferences = self.initialize_currency_preferences()
        self.conversion_rates = self.initialize_conversion_rates()  
        self.customers = []

    def initialize_currency_preferences(self):
        # Dictionary mapping countries to their preferred currencies
        currency_preferences = {
            Country.USA: Currency.USD,
            Country.UK: Currency.GBP,
            Country.CANADA: Currency.CAD,
            Country.GERMANY: Currency.EUR,
            Country.FRANCE: Currency.EUR,
            Country.AUSTRALIA: Currency.AUD,
            Country.JAPAN: Currency.JPY,
            Country.INDIA: Currency.INR,
            Country.CHINA: Currency.CNY
        }
        return currency_preferences

    def initialize_conversion_rates(self):
        # This is a simplified example. We would typically fetch conversion rates from an external API.
        conversion_rates = {
            Currency.USD: {Currency.USD: 1, Currency.GBP: 0.72, Currency.CAD: 1.24, Currency.EUR: 0.82,
                           Currency.AUD: 1.29, Currency.JPY: 109.45, Currency.INR: 73.1, Currency.CNY: 6.47},
            Currency.GBP: {Currency.USD: 1.39, Currency.GBP: 1, Currency.CAD: 1.73, Currency.EUR: 1.14,
                           Currency.AUD: 1.8, Currency.JPY: 151.64, Currency.INR: 100.92, Currency.CNY: 8.93},
            Currency.CAD: {Currency.USD: 0.81, Currency.GBP: 0.58, Currency.CAD: 1, Currency.EUR: 0.66,
                           Currency.AUD: 1.04, Currency.JPY: 87.73, Currency.INR: 58.31, Currency.CNY: 5.16},
            Currency.EUR: {Currency.USD: 1.22, Currency.GBP: 0.88, Currency.CAD: 1.52, Currency.EUR: 1,
                           Currency.AUD: 1.58, Currency.JPY: 132.82, Currency.INR: 88.29, Currency.CNY: 7.82},
            Currency.AUD: {Currency.USD: 0.78, Currency.GBP: 0.56, Currency.CAD: 0.96, Currency.EUR: 0.63,
                           Currency.AUD: 1, Currency.JPY: 84.1, Currency.INR: 55.94, Currency.CNY: 4.95},
            Currency.JPY: {Currency.USD: 0.0091, Currency.GBP: 0.0066, Currency.CAD: 0.0114, Currency.EUR: 0.0075,
                           Currency.AUD: 0.0119, Currency.JPY: 1, Currency.INR: 0.66, Currency.CNY: 0.058},
            Currency.INR: {Currency.USD: 0.014, Currency.GBP: 0.0099, Currency.CAD: 0.017, Currency.EUR: 0.011,
                           Currency.AUD: 0.018, Currency.JPY: 1.52, Currency.INR: 1, Currency.CNY: 0.088},
            Currency.CNY: {Currency.USD: 0.15, Currency.GBP: 0.11, Currency.CAD: 0.19, Currency.EUR: 0.13,
                           Currency.AUD: 0.20, Currency.JPY: 17.21, Currency.INR: 11.36, Currency.CNY: 1}
        }
        return conversion_rates

    def convert_amount(self, amount, source_currency, target_currency):
        if source_currency not in self.conversion_rates or target_currency not in self.conversion_rates:
            return None

        if source_currency not in self.operates_in or target_currency not in self.operates_in:
            return None

        conversion_rate = self.conversion_rates[source_currency][target_currency]
        converted_amount = amount * conversion_rate
        return converted_amount

    def get_preferred_currency(self, country):
        # Get the preferred currency for a given country
        return self.currency_preferences.get(country)


class BankServer:
    def __init__(self):
        self.users = {}  # Store user details {username: Customer}

    def process_request(self, command, username, password):
        if command == 'adduser':
            if username in self.users:
                return "User already exists\n"
        elif command == 'addBank':
            if username not in self.users or self.users[username].password != password:
                return "Invalid username or password\n"
            account_details = self.users[username].account_details
            return f"Welcome {username}, your account details are: {account_details}\n"
        else:
            return "Invalid command\n"

bank_server = BankServer()


app = Flask(__name__)
@app.route('/adduser', methods=['POST'])
def add_user():
    data = request.json
    username = data.get('username')
    password = data.get('password')
    if not username or not password:
        return jsonify({'error': 'Username and password are required'}), 400
    response = bank_server.process_request('adduser', username, password)
    return jsonify({'message': response})

@app.route('/addBank', methods=['POST'])
def sign_in():
    data = request.json
    username = data.get('username')
    password = data.get('password')
    bankname = data.get('bankname')
    amount = data.get('amount')
    if not username or not password:
        return jsonify({'error': 'Username and password are required'}), 400
    response = bank_server.process_request('addBank', username, password,bankname,amount)
    return jsonify({'message': response})





if __name__ == '__main__':
    americanbank = Bank(country=Country.USA, preferred_currency=Currency.USD, operates_in=[Country.USA, Country.INDIA])
    statebank = Bank(country=Country.INDIA, preferred_currency=Currency.USD, operates_in=[Country.USA, Country.INDIA]) 
    app.run(host='localhost', port=12345)
    
    
    
# Example usage:
# bank = Bank(country=Country.USA, preferred_currency=Currency.USD, operates_in=[Country.USA, Country.INDIA])

# # Convert 1000 USD to Indian Rupee (INR)
# amount = 1000
# source_currency = Currency.USD
# target_currency = Currency.INR
# converted_amount = bank.convert_amount(amount, source_currency, target_currency)
# print(f"{amount} {source_currency.value} is equivalent to {converted_amount} {target_currency.value}")

# # Get the preferred currency for India
# preferred_currency = bank.get_preferred_currency(Country.INDIA)
# print(f"The preferred currency for India is {preferred_currency.value}")