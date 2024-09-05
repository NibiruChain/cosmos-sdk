pub mod sync;

use cosmos_core_macros::service;

pub struct Context<Events=()> {
    _phantom: core::marker::PhantomData<Events>,
}

impl<Events> Context<Events> {
    pub fn self_address(&self) -> Address {
        todo!()
    }
    pub fn sender(&self) -> Address {
        todo!()
    }

    pub fn state_token(&self) -> StateToken {
        todo!()
    }

    pub fn emit_event<E>(&mut self, event: E) -> Result<()>
    where
        E: EventOf<Events>,
    {
        todo!()
    }

    pub fn new<NewEvents>(&self) -> Context<NewEvents> {
        todo!()
    }

    pub fn new_read(&self) -> ReadContext {
        todo!()
    }

    pub unsafe fn set_state_token(&mut self, state_token: StateToken) {
        todo!()
    }

    pub unsafe fn set_sender(&mut self, sender: Address) {
        todo!()
    }

    // consume_gas consumes gas from the gas meter.
    // It returns an error if the gas meter has run out.
    // This method uses interior mutability to update the gas meter so that it can
    // work from read-only references to the context.
    pub fn consume_gas(&self, gas: u64) -> Result<()> {
        todo!()
    }

    pub fn ok<T, E>(self, value: T) -> Response<T, E> {
        todo!()
    }

    pub fn err<T, E>(self, err: E) -> Response<T, E> {
        todo!()
    }
}

trait EventOf<Events> {}

impl<T> EventOf<T> for T {}

// impl<E1, E2> EventOf<(E1, E2)> for E1 {}
// impl<E1, E2> EventOf<(E1, E2)> for E2 {}
// impl<E1, E2, E3> EventOf<(E1, E2, E3)> for E1 {}
// impl<E1, E2, E3> EventOf<(E1, E2, E3)> for E2 {}
// impl<E1, E2, E3> EventOf<(E1, E2, E3)> for E3 {}
// impl<E1, E2, E3, E4> EventOf<(E1, E2, E3, E4)> for E1 {}
// impl<E1, E2, E3, E4> EventOf<(E1, E2, E3, E4)> for E2 {}
// impl<E1, E2, E3, E4> EventOf<(E1, E2, E3, E4)> for E3 {}
// impl<E1, E2, E3, E4> EventOf<(E1, E2, E3, E4)> for E4 {}
// impl<E1, E2, E3, E4, E5> EventOf<(E1, E2, E3, E4, E5)> for E1 {}
// impl<E1, E2, E3, E4, E5> EventOf<(E1, E2, E3, E4, E5)> for E2 {}
// impl<E1, E2, E3, E4, E5> EventOf<(E1, E2, E3, E4, E5)> for E3 {}
// impl<E1, E2, E3, E4, E5> EventOf<(E1, E2, E3, E4, E5)> for E4 {}
// impl<E1, E2, E3, E4, E5> EventOf<(E1, E2, E3, E4, E5)> for E5 {}

pub struct ReadContext {}

impl ReadContext {
    pub fn new(&self) -> ReadContext {
        todo!()
    }

    pub fn new_read(&self) -> ReadContext {
        todo!()
    }
}

impl ReadContext {
    pub fn consume_gas(&self, gas: u64) -> Result<()> {
        todo!()
    }

    pub fn self_address(&self) -> Address {
        todo!()
    }
    pub fn sender(&self) -> Address {
        todo!()
    }

    pub fn state_token(&self) -> StateToken {
        todo!()
    }
}

pub type StateToken = [u8; 32];

pub type Result<T> = core::result::Result<T, String>;

pub struct MessagePacket {
    header: MessagePacketHeader, // 64 + 64 + 32 + 8 + 128 + 64 + 664 = 1024
    data: [u8; 0xFC00], // 64512
}

pub struct MessagePacketHeader {
    address: Address, // 64
    caller: Address, // 64
    state_token: [u8; 32], //32
    gas_limit: u64, //8
    message_name: MessageName, //128
    params: [BufferRef; 4], // 16 * 4 = 64
    padding: [u8; 664], // 664
}

#[derive(Clone, Copy, PartialEq, Eq)]
pub struct Address {
    len: u8,
    bytes: [u8; 63],
}

impl Default for Address {
    fn default() -> Self {
        Self {
            len: 0,
            bytes: [0; 63],
        }
    }
}

pub struct MessageName {
    len: u8,
    bytes: [u8; 127],
}

pub struct BufferRef {
    pointer: u64,
    capacity: u32,
    len: u32,
}

#[derive(Default, Copy, Clone, PartialEq, Eq, Ord, PartialOrd)]
pub struct Time(u64);

up#[service]
pub trait BlockService {
    fn current_time(&self, ctx: &Context) -> crate::Result<Time>;
}

pub trait Account {}

pub trait Module {}

pub trait Serializable {}

pub struct Item<T> {
    _phantom: core::marker::PhantomData<T>,
}

impl<T: Default> Item<T> {
    pub fn get(&self, ctx: &Context) -> Result<T> {
        todo!()
    }

    pub fn set(&self, ctx: &mut Context, value: &T) -> Result<()> {
        todo!()
    }
}

pub struct Map<K, V> {
    _phantom: core::marker::PhantomData<(K, V)>,
}

impl<K, V> Map<K, V> {
    pub fn get(&self, ctx: ReadContext, key: &K) -> Result<Option<V>> {
        todo!()
    }

    pub fn set(&self, ctx: Context, key: &K, value: &V) -> Result<()> {
        todo!()
    }
}


pub trait Table {}

pub trait OnCreate {
    type InitMessage;

    fn on_create(&self, ctx: &mut Context, msg: &Self::InitMessage) -> Result<()>;
}

pub trait State {}

pub trait Service {
    fn client() -> Box<Self>;
    fn client_with_ctx<F>(ctx_fn: F) -> Box<Self>
    where
        F: FnOnce(&mut Context);
}

pub type Response<T, E=String> = core::result::Result<ResponseBody<T>, E>;

impl <T> ResponseBody<T> {
    pub fn read(&self) -> core::result::Result<&T, String> {
        todo!()
    }
}

pub struct ResponseBody<T> {}