# -*- coding: utf-8 -*-

"""
    pretty print json
    from: <http://stackoverflow.com/questions/12943819/how-to-python-prettyprint-a-json-file>
"""

import json

def pretty_json(json_str):
    parsed = json.loads(json_str)
    return json.dumps(parsed, indent=4, sort_keys=True)

if __name__ == '__main__':
    your_json = '["foo", {"bar":["baz", null, 1.0, 2]}]'
    print pretty_json(your_json)

