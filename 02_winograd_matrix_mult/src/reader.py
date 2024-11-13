from typing import Callable

def get_value(msg: str = "", default: int = None, check: Callable[[int], bool] = None) -> int:
    val = None 
    if default is not None:
        msg += f" (Press enter for default value = {default}): "
    while val is None:
        try: 
            val = int(input(msg))
            if not (check is None or check(val)):
                print("[ERR]: Wrong number. Try again.")
                val = None
        except ValueError:
            if default is not None:
                val = default
                break
            print("[ERR]: Invalid data. Try again.")

    return val
