
EXPECTED_BAKE_TIME = 40


def bake_time_remaining(time_in):
    """Calculate the bake time remaining.

    :param time_in: int - the amount of time the lasagna has been in the oven.
    :return: int - the time the lasagna still has to go in the oven.

    This function calculates how much bake time is left for the lasagna.
    """
    
    return EXPECTED_BAKE_TIME - time_in


def preparation_time_in_minutes(number_of_layers):
    """Calculate the preparation time.

    :param number_of_layers: int - the number of layers in the lasagna.
    :return: int - total time to prepare the lasagna.

    This function takes the number of layers in the lasagna and calculates how much 
    time it takes to prepare before baking.
    """
    
    return number_of_layers * 2



def elapsed_time_in_minutes(number_of_layers, elapsed_bake_time):
    """"Calculate the elapsed cooking time.

    :param number_of_layers: int - the number of layers in the lasagna.
    :param elapsed_bake_time: int - total time elapsed (in minutes) preparing and cooking.
    :return int

    This function takes two integers representing the number of layers of lasagna and the
    time already spent baking and calculates the total elapsed minutes spent making lasagna.
    """
    
    return preparation_time_in_minutes(number_of_layers) + elapsed_bake_time
