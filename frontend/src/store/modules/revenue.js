import axios from 'axios';
export default {
  namespaced: true,
  state: {
    data: {
      title: 'Revenues',
      columns: ['name', 'type', 'description', 'datetime', 'amount'],
      data: [],
      total: 0
    },
    isLoading: false,
    filters: {
      page: 1,
      size: 10
    },
    options: {
      groups: [
        { value: null, text: 'Select a group'},
        { value: '1', text: 'group 1'},
        { value: '2', text: 'group 2'},
      ],
      types: [
        { value: null, text: 'Please select a type' },
        { value: 'Продукты', text: 'Продукты' },
        { value: 'Депозит', text: 'Депозит' },
        { value: 'Кредит', text: 'Кредит' },
        { value: 'Рассрочка', text: 'Рассрочка' },
        { value: 'Фастфуд', text: 'Фастфуд' },
        { value: 'Доставка еды', text: 'Доставка еды' },
        { value: 'Оплата интернета', text: 'Оплата интернета' },
        { value: 'Оплата тарифного плана', text: 'Оплата тарифного плана' },
        { value: 'Одежда', text: 'Одежда' },
        { value: 'Электронная техника', text: 'Электронная техника' },
        { value: 'Цветы', text: 'Цветы' },
        { value: 'Подарки', text: 'Подарки' },
        { value: 'Кинотеатр', text: 'Кинотеатр' },
        { value: 'Авиабилеты', text: 'Авиабилеты' },
        { value: 'ЖД билеты', text: 'ЖД билеты' },
        { value: 'Такси', text: 'Такси' },
        { value: 'Украшения', text: 'Украшения' },
        { value: 'Компьютерные игры', text: 'Компьютерные игры' },
        { value: 'Другие развлечения', text: 'Другие развлечения' },
        { value: 'Кафе/рестораны', text: 'Кафе/рестораны' },
        { value: 'Спорт', text: 'Спорт' },
        { value: 'Книги', text: 'Книги' },
        { value: 'Здоровье', text: 'Здоровье' },
        { value: 'Отели', text: 'Отели' },
        { value: 'Аренда жилья', text: 'Аренда жилья' },
        { value: 'Перевод', text: 'Перевод' },
        { value: 'Налоги', text: 'Налоги' },
        { value: 'Вывод денег', text: 'Вывод денег' },
        { value: 'Пополнение счета', text: 'Пополнение счета' },
        { value: 'Снятия с депозита', text: 'Снятия с депозита' },
        { value: 'Получение перевода', text: 'Получение перевода' },
        { value: 'Долг', text: 'Долг' },
      ]
    },
  },
  actions: {
    async getRevenues({ state, commit }){
     try {
      const r = await axios.get(`revenue?page=${state.filters.page}&size=${state.filters.size}`)
      commit('SET_REVENUES', r.data.payload)
     } catch (error) {
      console.log(error);
     }
    },
    async addRevenue({ commit }, payload){
      try {
        await axios.post('revenue', payload)
      } catch (error) {
        console.log(error);
      }

    }
  },
  mutations: {
    SET_REVENUES(state, payload){
      state.data.data = payload.revenues
      state.data.total = payload.total
    },
  },
}