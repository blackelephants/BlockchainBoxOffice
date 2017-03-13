import json

import requests
from django.shortcuts import render

user_type = "user_type1_1"
end_point = 'https://f78040cda74f4ab5bfe4e3c5c53a9a4b-vp1.us.blockchain.ibm.com:5001/chaincode'
chaincode_name = "327771fb6f34cc8fb0f55aa10695aa49eefcc0e73414bb0606b2a9fd6bdb495c14fb2940de5075db751723cf758dd93a3d218c58b88faac481838060cdbac76f"


def request_bluemix(end_point, chaincode_name, function_name, args_array, method='invoke'):
    payload = {
        "jsonrpc": "2.0",
        "method": method,  # invoke or query
        "params": {
            "type": 1,
            "chaincodeID": {
                "name": chaincode_name
            },
            "ctorMsg": {
                "function": function_name,
                "args": args_array
            },
            "secureContext": user_type
        },
        "id": 0
    }
    headers = {'Content-type': 'application/json', 'Accept': 'text/plain', 'content-type': 'application/json'}

    r = requests.post(end_point, data=json.dumps(payload), headers=headers)
    print(r)
    print(r.text)
    return r.json()


if __name__ == '__main__':
    function_name = "registerCinema"
    args_array = [
        "影院名称",
        "所属院线"
    ]
    request_bluemix(end_point, chaincode_name, function_name, args_array)


def index(request):
    """A view of all movies."""

    return render(request, 'index.html', {})


def cinemas(request):
    return render(request, 'cinemas.html', {"cinemas": ''})


def cinema_reg(request):
    function_name = "registerCinema"
    args_array = [
        "杭州西湖万达电影院",
        "万达电影"
    ]
    cinema_reg = request_bluemix(end_point, chaincode_name, function_name, args_array)
    return render(request, 'cinemas.html', {'cinema_reg': cinema_reg})


def cinema_query(request):
    function_name = "queryCinema"
    args_array = [
        "杭州西湖万达电影院"
    ]
    cinema_query = request_bluemix(end_point, chaincode_name, function_name, args_array, method='query')

    return render(request, 'cinemas.html', {"cinema_query": json.loads(cinema_query['result']['message'])})


def halls(request):
    return render(request, 'halls.html', {})


def hall_reg(request):
    function_name = "registerVideoHall"
    args_array = [
        "金沙巨幕1号厅",
        "杭州西湖万达电影院",
        "10",
        "15"
    ]
    hall_reg = request_bluemix(end_point, chaincode_name, function_name, args_array)
    return render(request, 'halls.html', {'hall_reg': hall_reg})


def hall_query(request):
    function_name = "queryVideoHall"
    args_array = [
        "金沙巨幕1号厅",
        "杭州西湖万达电影院"
    ]
    hall_query = request_bluemix(end_point, chaincode_name, function_name, args_array, method='query')
    return render(request, 'halls.html', {'hall_query': json.loads(hall_query['result']['message'])})


def seller(request):
    return render(request, 'seller.html', {})


def audience(request):
    return render(request, 'audience.html', {})


def capital(request):
    return render(request, 'capital.html', {})

def capital_query(request):
    function_name = "clear"
    args_array = [
        "金刚狼3：殊死一战"
    ]
    capital_query = request_bluemix(end_point, chaincode_name, function_name, args_array, method='query')
    return render(request, 'capital.html', {'capital_query': json.loads(capital_query['result']['message'])})


def schedule(request):
    function_name = "planMovie"
    args_array = [
        "plan1",
        "生化危机：终章",
        "杭州西湖万达电影院",
        "VIP巨幕2号厅",
        "2017-03-13 9:00",
        "2017-03-13 11:00",
        "2017-03-13 15:00"
    ]
    plan = request_bluemix(end_point, chaincode_name, function_name, args_array, method='invoke')
    return render(request, 'halls.html', {'plan': plan})


def schedule_query(request):
    function_name = "queryPlan"
    args_array = [
        "plan1"
    ]
    plan_query = request_bluemix(end_point, chaincode_name, function_name, args_array, method='query')
    return render(request, 'halls.html', {'plan_query': json.loads(plan_query['result']['message'])})


def all_plan_query(request):
    function_name = "queryAllPlan"
    args_array = [
    ]
    all_plan_query = request_bluemix(end_point, chaincode_name, function_name, args_array, method='query')
    return render(request, 'audience.html', {'all_plan_query':  json.loads(all_plan_query['result']['message'])})


def lock_ticket(request):
    import random
    function_name = "lockTicket"
    args_array = [
        "plan1:"+str(random.randint(0,10))+"-"+str(random.randint(0,10)),
        "50"
    ]
    lock_ticket = request_bluemix(end_point, chaincode_name, function_name, args_array, method='invoke')
    return render(request, 'audience.html', {'lock_ticket': lock_ticket})


def check_ticket(request):
    function_name = "checkTicket"
    args_array = [
        "plan1:0-0"
    ]
    check_ticket = request_bluemix(end_point, chaincode_name, function_name, args_array, method='invoke')
    return render(request, 'audience.html', {'check_ticket': check_ticket})


