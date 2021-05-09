package L3.Z4;

class Figury {

    enum OneArg {
        CIRCLE,
        SQUARE,
        PENTAGON
    }

    public double OneArgObliczPole (double dupa, OneArg type) {

        switch(type)
        {
        case CIRCLE:
            return dupa * dupa * Math.PI;
        case PENTAGON:
            return dupa;
        case SQUARE:
            return dupa *dupa;
        default:
            return dupa;
            
        }
    }
}